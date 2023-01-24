import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/bloc/control_bloc.dart';
import 'package:smiot/bloc/display_bloc.dart';
import 'package:smiot/widgets/button.dart';
import 'package:smiot/widgets/circular_percent_display.dart';
import 'package:smiot/widgets/line_chart_display.dart';
import 'package:smiot/widgets/number_button.dart';
import 'package:smiot/widgets/number_display.dart';
import 'package:smiot/widgets/on_off_display.dart';
import 'package:smiot/widgets/slider_control.dart';
import 'package:smiot/widgets/splash.dart';
import 'package:smiot/widgets/switch.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class DashboardScreen extends StatefulWidget {
  final String boxId;
  final String canSub;
  DashboardScreen({
    Key? key,
    required this.boxId,
    required this.canSub,
  }) : super(key: key);

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

final displayWidget = [
  LineChartDisplay,
  CircularPercent,
  NumberDisplay,
];

class _DashboardScreenState extends State<DashboardScreen> {
  late IO.Socket _socket;
  var connectState = false;

  _connectSocket() {
    _socket.io
      ..disconnect()
      ..connect();
    _socket.onConnect((data) {
      if (mounted) {
        setState(() {
          connectState = true;
        });
      }
    });
    _socket.onConnectError((data) {
      if (mounted) {
        _socket.io
          ..disconnect()
          ..connect();
        setState(() {
          connectState = false;
        });
      }
    });
    _socket.onDisconnect((data) {
      if (mounted) {
        _socket.io
          ..disconnect()
          ..connect();
        setState(() {
          connectState = false;
        });
      }
    });
  }

  @override
  void initState() {
    super.initState();
    _socket = IO.io(
      'wss://api.rocket-translate.com/',
      IO.OptionBuilder().setTransports(['websocket']).build(),
    );
    _connectSocket();
    context.read<DisplayBloc>().add(GetDisplayEvent(widget.boxId));
    context.read<ControlBloc>().add(GetControlEvent(widget.boxId));
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Dashboard ${connectState ? '🔴' : '⚫️'}',
          style:
              TextStyle(color: Colors.grey[800], fontWeight: FontWeight.bold),
        ),
        iconTheme: const IconThemeData(color: Colors.black),
        backgroundColor: Colors.white,
      ),
      body: Container(
        color: Colors.grey.shade100,
        padding: const EdgeInsets.all(10),
        child: ListView(children: [
          BlocBuilder<DisplayBloc, MyState>(
            builder: (context, state) {
              if (state is GetDisplaysStateSuccess) {
                return Column(children: [
                  for (var d in state.displays)
                    d.widget.name == 'DLineChart'
                        ? LineChartDisplay(
                            canSub: widget.canSub,
                            display: d,
                            socket: _socket,
                          )
                        : d.widget.name == 'DCircularPercent'
                            ? CircularPercent(
                                canSub: widget.canSub,
                                display: d,
                                socket: _socket,
                              )
                            : d.widget.name == 'DNumber'
                                ? NumberDisplay(
                                    canSub: widget.canSub,
                                    display: d,
                                    socket: _socket,
                                  )
                                : d.widget.name == 'DOnOff'
                                    ? OnOffDisplay(
                                        canSub: widget.canSub,
                                        display: d,
                                        socket: _socket,
                                      )
                                    : d.widget.name == 'DOnOffSwitch'
                                        ? OnOffDisplay(
                                            canSub: widget.canSub,
                                            display: d,
                                            socket: _socket,
                                          )
                                        : const SizedBox()
                ]);
              }
              if (state is StateLoading) {
                return const Splash();
              } else {
                return const SizedBox();
              }
            },
          ),
          BlocBuilder<ControlBloc, MyState>(
            builder: (context, state) {
              if (state is GetControlStateSuccess) {
                return Column(children: [
                  for (var c in state.controls)
                    c.widget.name == 'CButton'
                        ? Button(
                            canSub: widget.canSub,
                            controls: c,
                            socket: _socket,
                          )
                        : c.widget.name == 'CSlider'
                            ? SliderControl(
                                canSub: widget.canSub,
                                controls: c,
                                socket: _socket,
                              )
                            : c.widget.name == 'CSwitch'
                                ? SwitchButton(
                                    canSub: widget.canSub,
                                    controls: c,
                                    socket: _socket,
                                  )
                                : c.widget.name == 'CButtonNumber'
                                    ? NumberButton(
                                        canSub: widget.canSub,
                                        controls: c,
                                        socket: _socket,
                                      )
                                    : const SizedBox()
                ]);
              }
              if (state is StateLoading) {
                return const Splash();
              } else {
                return const SizedBox();
              }
            },
          ),
        ]),
      ),
    );
  }
}
