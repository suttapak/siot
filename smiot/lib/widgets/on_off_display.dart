import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/models/display_data_models.dart';
import 'package:smiot/models/display_models.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

import 'package:http/http.dart' as http;

const _storage = FlutterSecureStorage();

Future<List<DisplayDatum>> fetchData(String boxId, int displayId) async {
  final token = await _storage.read(key: 'accessToken');

  final response = await http.get(
      Uri.parse(
          'https://api.rocket-translate.com/boxes/$boxId/displays/$displayId/data'),
      headers: <String, String>{
        'context-type': 'application/json',
        'Authorization': 'Bearer $token'
      });

  if (response.statusCode == 200) {
    // If the server did return a 200 OK response,
    // then parse the JSON.
    final json = jsonDecode(response.body) as List;
    return json.map((e) => DisplayDatum.fromJson(e)).toList();
  } else {
    // If the server did not return a 200 OK response,
    // then throw an exception.
    throw Exception('Failed to load album');
  }
}

class OnOffDisplay extends StatefulWidget {
  const OnOffDisplay({
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
    return 'DNumber';
  }

  @override
  State<OnOffDisplay> createState() => _OnOffDisplayState();
}

class _OnOffDisplayState extends State<OnOffDisplay> {
  late List<DisplayDatum> data;
  late final IO.Socket _socket;

  @override
  void initState() {
    super.initState();
    _requestData();
    data = widget.display.displayData;
    _socket = widget.socket;
    _subScript();
  }

  void _requestData() async {
    if (mounted) {
      final d = await fetchData(widget.display.boxId, widget.display.id);
      setState(() {
        data = d;
      });
    }
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
            child: Text(
              data.isEmpty
                  ? 'N/A'
                  : data[data.length - 1].data == 1
                      ? 'On'
                      : 'Off',
              style: const TextStyle(
                fontSize: 22,
                fontWeight: FontWeight.w500,
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
