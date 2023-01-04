import 'package:flutter/material.dart';
import 'package:smiot/models/control_data_models.dart';
import 'package:smiot/models/control_models.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class Button extends StatefulWidget {
  const Button({
    Key? key,
    required this.controls,
    required this.socket,
    required this.canSub,
  }) : super(key: key);

  final Controls controls;
  final IO.Socket socket;
  final String canSub;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'CButton';
  }

  @override
  State<Button> createState() => _ButtonState();
}

class _ButtonState extends State<Button> {
  late List<ControlDatum> data;
  late final IO.Socket _socket;

  @override
  void initState() {
    super.initState();

    data = widget.controls.controlData;
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
    super.dispose();
  }

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
              borderRadius: BorderRadius.circular(15),
              color: Colors.white,
            ),
            child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: data.isEmpty
                      ? Colors.grey[300]
                      : data[data.length - 1].data % 2 != 0
                          ? Colors.yellow[300]
                          : Colors.grey[300],
                  minimumSize: Size(110, 50),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(30), // <-- Radius
                  ),
                ),
                child: Text(
                  data.isEmpty
                      ? 'NULL'
                      : data[data.length - 1].data % 2 != 0
                          ? 'On'
                          : 'Off',
                  style: TextStyle(fontSize: 22, color: Colors.grey[800]),
                ),
                onPressed: () => _setMessage(data.isEmpty
                    ? 1
                    : data[data.length - 1].data % 2 != 0
                        ? 0
                        : 1)),
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
}
