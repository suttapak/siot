import 'package:flutter/material.dart';

import 'package:smiot/models/control_data_models.dart';
import 'package:smiot/models/control_models.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SwitchButton extends StatefulWidget {
  const SwitchButton(
      {super.key,
      required this.controls,
      required this.socket,
      required this.canSub});

  final Controls controls;
  final IO.Socket socket;
  final String canSub;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'CSwitch';
  }

  @override
  State<SwitchButton> createState() => _SwitchButtonState();
}

class _SwitchButtonState extends State<SwitchButton> {
  final animationDuration = const Duration(milliseconds: 100);
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
            child: GestureDetector(
              onTap: () {
                _setMessage(data.isEmpty
                    ? 1
                    : data[data.length - 1].data % 2 != 0
                        ? 0
                        : 1);
              },
              child: AnimatedContainer(
                height: 55,
                width: 110,
                duration: animationDuration,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(30),
                  color: Colors.grey.shade200,
                  border: Border.all(color: Colors.white, width: 2),
                ),
                child: AnimatedAlign(
                  duration: animationDuration,
                  alignment: data.isEmpty
                      ? Alignment.centerLeft
                      : data[data.length - 1].data % 2 != 0
                          ? Alignment.centerRight
                          : Alignment.centerLeft,
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 2),
                    child: Container(
                      width: 45,
                      height: 45,
                      decoration: const BoxDecoration(
                        shape: BoxShape.circle,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
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
}
