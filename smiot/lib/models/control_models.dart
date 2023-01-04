// To parse this JSON data, do
//
//     final controls = controlsFromJson(jsonString);

import 'package:meta/meta.dart';
import 'package:smiot/models/control_data_models.dart';
import 'package:smiot/models/layout_models.dart';
import 'dart:convert';

import 'package:smiot/models/widget_models.dart';

List<Controls> controlsFromJson(String str) =>
    List<Controls>.from(json.decode(str).map((x) => Controls.fromJson(x)));

String controlsToJson(List<Controls> data) =>
    json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class Controls {
  Controls({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.key,
    required this.description,
    required this.boxId,
    required this.layoutId,
    required this.widgetId,
    required this.controlData,
    required this.widget,
    required this.layout,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String name;
  String key;
  String description;
  String boxId;
  int layoutId;
  int widgetId;
  List<ControlDatum> controlData;
  Widget widget;
  Layout layout;

  factory Controls.fromJson(Map<String, dynamic> json) => Controls(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        key: json["key"],
        description: json["description"],
        boxId: json["BoxId"],
        layoutId: json["layoutId"],
        widgetId: json["widgetId"],
        controlData: List<ControlDatum>.from(
            json["controlData"].map((x) => ControlDatum.fromJson(x))),
        widget: Widget.fromJson(json["widget"]),
        layout: Layout.fromJson(json["layout"]),
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "name": name,
        "key": key,
        "description": description,
        "BoxId": boxId,
        "layoutId": layoutId,
        "widgetId": widgetId,
        "controlData": List<dynamic>.from(controlData.map((x) => x.toJson())),
        "widget": widget.toJson(),
        "layout": layout.toJson(),
      };
}
