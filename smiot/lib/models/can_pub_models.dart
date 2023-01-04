class CanPub {
  CanPub({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.canPublish,
    required this.boxId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String canPublish;
  String boxId;

  factory CanPub.fromJson(Map<String, dynamic> json) => CanPub(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        canPublish: json["canPublish"],
        boxId: json["BoxId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "canPublish": canPublish,
        "BoxId": boxId,
      };
}
