import 'package:flutter/material.dart';

import 'package:smiot/models/control_data_models.dart';
import 'package:smiot/models/control_models.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SliderControl extends StatefulWidget {
  const SliderControl({
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
    return 'CSlider';
  }

  @override
  State<SliderControl> createState() => _SliderControlState();
}

class _SliderControlState extends State<SliderControl> {
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
            child: Column(
              children: [
                Center(
                  child: Text(
                      '${data.isEmpty ? 0 : data[data.length - 1].data > 100 ? 100 : data[data.length - 1].data}'),
                ),
                Slider(
                  value: data.isEmpty
                      ? 0
                      : data[data.length - 1].data > 100
                          ? 100.0
                          : data[data.length - 1].data.toDouble(),
                  max: 100,
                  label: data.isEmpty
                      ? '0'
                      : data[data.length - 1].data > 100
                          ? 100.0.toString()
                          : data[data.length - 1].data.toString(),
                  onChanged: (value) => value,
                  onChangeEnd: ((value) => _setMessage(value.toInt())),
                ),
              ],
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
