import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:intl/intl.dart';
import 'package:flutter/material.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/bloc/display_data_bloc.dart';
import 'package:smiot/models/display_data_models.dart';
import 'package:smiot/models/display_models.dart';
import 'package:syncfusion_flutter_charts/charts.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class LineChartDisplay extends StatefulWidget {
  // ignore: prefer_const_constructors_in_immutables
  LineChartDisplay({
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
    return 'DLineChart';
  }

  @override
  State<LineChartDisplay> createState() => _LineChartDisplayState();
}

class _LineChartDisplayState extends State<LineChartDisplay> {
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
            alignment: Alignment.center,
            width: double.infinity,
            height: 200,
            padding: const EdgeInsets.all(10),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(15),
              color: Colors.white,
            ),
            child:
                //Initialize the chart widget
                BlocListener<DisplayDataBloc, MyState>(
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
              child: SfCartesianChart(
                primaryXAxis: CategoryAxis(),
                tooltipBehavior: TooltipBehavior(enable: true),
                series: <ChartSeries<DisplayDatum, String>>[
                  LineSeries<DisplayDatum, String>(
                      dataSource: data,
                      xValueMapper: (DisplayDatum datum, i) =>
                          DateFormat('hh:mm:s a dd').format(datum.updatedAt),
                      yValueMapper: (DisplayDatum datum, _) => datum.data,
                      name: 'Sales',
                      // Enable data label
                      dataLabelSettings:
                          const DataLabelSettings(isVisible: true))
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
