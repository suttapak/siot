class ControlDatum {
  ControlDatum({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.data,
    required this.label,
    required this.controlId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  int data;
  String label;
  int controlId;

  factory ControlDatum.fromJson(Map<String, dynamic> json) => ControlDatum(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        data: json["data"],
        label: json["label"],
        controlId: json["controlId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "data": data,
        "label": label,
        "controlId": controlId,
      };
}
