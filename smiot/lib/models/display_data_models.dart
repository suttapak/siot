class DisplayDatum {
  DisplayDatum({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.data,
    required this.label,
    required this.displayId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  num data;
  String label;
  int displayId;

  factory DisplayDatum.fromJson(Map<String, dynamic> json) => DisplayDatum(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        data: json["data"],
        label: json["label"],
        displayId: json["displayId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "data": data,
        "label": label,
        "displayId": displayId,
      };
}
