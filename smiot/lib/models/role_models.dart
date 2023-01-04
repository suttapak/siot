class Role {
  Role({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.permissionState,
    required this.displayName,
    required this.description,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String name;
  int permissionState;
  String displayName;
  String description;

  factory Role.fromJson(Map<String, dynamic> json) => Role(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        permissionState: json["permissionState"],
        displayName: json["displayName"],
        description: json["description"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "name": name,
        "permissionState": permissionState,
        "displayName": displayName,
        "description": description,
      };
}
