class BoxSecret {
  BoxSecret({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.secret,
    required this.boxId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String secret;
  String boxId;

  factory BoxSecret.fromJson(Map<String, dynamic> json) => BoxSecret(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        secret: json["secret"],
        boxId: json["BoxId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "secret": secret,
        "BoxId": boxId,
      };
}
