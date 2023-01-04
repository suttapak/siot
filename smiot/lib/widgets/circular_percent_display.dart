import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:percent_indicator/percent_indicator.dart';

import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/bloc/display_data_bloc.dart';
import 'package:smiot/models/display_data_models.dart';
import 'package:smiot/models/display_models.dart';

import 'package:socket_io_client/socket_io_client.dart' as IO;

class CircularPercent extends StatefulWidget {
  const CircularPercent({
    Key? key,
    required this.display,
    required this.socket,
    required this.canSub,
  }) : super(key: key);

  final Displays display;
  final IO.Socket socket;
  final String canSub;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'DCircularPercent';
  }

  @override
  State<CircularPercent> createState() => _CircularPercentState();
}

class _CircularPercentState extends State<CircularPercent> {
  late List<DisplayDatum> data;
  late final IO.Socket _socket;

  @override
  void initState() {
    super.initState();
    context.read<DisplayDataBloc>().add(
          GetDisplayDataEvent(
            displayId: widget.display.id,
            boxId: widget.display.boxId,
          ),
        );
    data = widget.display.displayData;
    _socket = widget.socket;
    _subScript();
  }

  void _subScript() {
    _socket.emit('subscript', {
      'boxId': widget.display.boxId,
      'key': '${widget.canSub}/${widget.display.key}'
    });
    _socket.on('${widget.canSub}/${widget.display.key}', _getMessage);
  }

  void _getMessage(dynamic value) {
    final json = value['displayData'] as List;
    var dataDisplay = json.map((e) => DisplayDatum.fromJson(e)).toList();
    if (mounted) {
      setState(() {
        data = dataDisplay;
        data.sort(((a, b) => a.id.compareTo(b.id)));
      });
    }
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
        alignment: const Alignment(.98, -.98),
        children: [
          Container(
            alignment: Alignment.topCenter,
            width: double.infinity,
            height: 200,
            padding: const EdgeInsets.all(10),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(15),
              color: Colors.white,
            ),
            child: BlocListener<DisplayDataBloc, MyState>(
              listener: (context, state) {
                if (state is GetDisplayDataStateSuccess) {
                  if (state.displayData.isEmpty) return;
                  if (state.displayData[0].displayId != widget.display.id) {
                    return;
                  }
                  setState(() {
                    data = state.displayData;
                    data.sort(((a, b) => a.id.compareTo(b.id)));
                  });
                }
              },
              child: Column(
                children: [
                  CircularPercentIndicator(
                    circularStrokeCap: CircularStrokeCap.round,
                    radius: 90.0,
                    lineWidth: 12.0,
                    percent: data.isEmpty
                        ? 0
                        : ((data[data.length - 1].data % 100) * pow(10, -2))
                            .toDouble(),
                    animation: true,
                    center: Text(
                      '${data.isEmpty ? 0 : ((data[data.length - 1].data % 100))}%',
                      style: const TextStyle(
                          fontSize: 18, fontWeight: FontWeight.w500),
                    ),
                    animationDuration: 1200,
                    progressColor: Colors.blue,
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
            child: Text(widget.display.key),
          )
        ],
      ),
    );
  }
}
