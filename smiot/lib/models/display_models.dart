import 'package:smiot/models/display_data_models.dart';
import 'package:smiot/models/layout_models.dart';
import 'package:smiot/models/widget_models.dart';

class Displays {
  Displays({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.key,
    required this.description,
    required this.boxId,
    required this.layoutId,
    required this.widgetId,
    required this.displayData,
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
  List<DisplayDatum> displayData;
  Widget widget;
  Layout layout;

  factory Displays.fromJson(Map<String, dynamic> json) => Displays(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        key: json["key"],
        description: json["description"],
        boxId: json["BoxId"],
        layoutId: json["layoutId"],
        widgetId: json["widgetId"],
        displayData: List<DisplayDatum>.from(
            json["displayData"].map((x) => DisplayDatum.fromJson(x))),
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
        "displayData": List<dynamic>.from(displayData.map((x) => x.toJson())),
        "widget": widget.toJson(),
        "layout": layout.toJson(),
      };
}
