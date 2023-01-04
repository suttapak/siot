class Widget {
  Widget({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.description,
    required this.dataType,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String name;
  String description;
  String dataType;

  factory Widget.fromJson(Map<String, dynamic> json) => Widget(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        description: json["description"],
        dataType: json["dataType"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "name": name,
        "description": description,
        "dataType": dataType,
      };
}
