import 'package:flutter/services.dart';

import 'package:flutter/material.dart';
import 'package:smiot/models/control_data_models.dart';
import 'package:smiot/models/control_models.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class NumberButton extends StatefulWidget {
  const NumberButton({
    Key? key,
    required this.controls,
    required this.socket,
    required this.canSub,
  }) : super(key: key);

  final Controls controls;
  final IO.Socket socket;
  final String canSub;

  @override
  State<NumberButton> createState() => _NumberButtonState();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'CButtonNumber';
  }
}

class _NumberButtonState extends State<NumberButton> {
  late List<ControlDatum> data;
  late final IO.Socket _socket;

  @override
  void initState() {
    super.initState();

    data = widget.controls.controlData;
    statusController.text =
        data.isEmpty ? '0' : data[data.length - 1].data.toString();

    _socket = widget.socket;
    _subScript();
  }

  void _subScript() {
    _socket.emit('subscript', {
      'boxId': widget.controls.boxId,
      'key': '${widget.canSub}/${widget.controls.key}'
    });
    _socket.on('${widget.canSub}/${widget.controls.key}', _getMessage);
  }

  void _getMessage(dynamic value) {
    final json = value['controlData'] as List;
    var controlData = json.map((e) => ControlDatum.fromJson(e)).toList();
    if (mounted) {
      setState(() {
        data = controlData;
        data.sort(((a, b) => a.id.compareTo(b.id)));
        statusController.text =
            data.isEmpty ? '0' : data[data.length - 1].data.toString();
        statusController.selection = TextSelection.fromPosition(
            TextPosition(offset: statusController.text.length));
      });
    }
  }

  void _setMessage(int value) {
    var msg = <String, dynamic>{
      'data': value,
      'boxId': widget.controls.boxId,
      'key': widget.controls.key,
    };
    _socket.emit('publish', msg);
  }

  @override
  void dispose() {
    statusController.dispose();

    super.dispose();
  }

  var statusController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Stack(
        alignment: const Alignment(.98, -.9),
        children: [
          Container(
            alignment: Alignment.center,
            width: double.infinity,
            height: 100,
            padding: const EdgeInsets.all(10),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
              color: Colors.white,
            ),
            child: Padding(
              padding: const EdgeInsets.only(left: 30, right: 30),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Container(
                    decoration: BoxDecoration(
                      color: Colors.grey.shade400,
                      borderRadius: BorderRadius.circular(15),
                    ),
                    child: IconButton(
                      icon: const Icon(
                        Icons.remove,
                        size: 30,
                      ),
                      onPressed: () {
                        _setMessage(
                            data.isEmpty ? 0 : data[data.length - 1].data - 1);
                      },
                    ),
                  ),
                  buildForm(),
                  Container(
                    decoration: BoxDecoration(
                      color: Colors.grey.shade400,
                      borderRadius: BorderRadius.circular(15),
                    ),
                    child: IconButton(
                      icon: const Icon(
                        Icons.add,
                        size: 30,
                      ),
                      onPressed: () {
                        _setMessage(
                            data.isEmpty ? 1 : data[data.length - 1].data + 1);
                      },
                    ),
                  ),
                ],
              ),
            ),
          ),
          Container(
            decoration: BoxDecoration(
              color: Colors.grey.shade300,
              borderRadius: BorderRadius.circular(6),
            ),
            padding:
                const EdgeInsets.only(left: 10, top: 4, bottom: 4, right: 10),
            child: Text(widget.controls.key),
          )
        ],
      ),
    );
  }

  Widget buildForm() => SizedBox(
      width: 200,
      child: TextField(
        textAlign: TextAlign.center,
        style: const TextStyle(fontSize: 22),
        controller: statusController,
        inputFormatters: [
          FilteringTextInputFormatter.allow(RegExp(r'[0-9]')),
          FilteringTextInputFormatter.digitsOnly
        ],
        keyboardType: TextInputType.number,
        onChanged: ((value) {
          _setMessage(int.parse(value));
        }),
      ));
}
