class CanSub {
  CanSub({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.canSubscribe,
    required this.boxId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String canSubscribe;
  String boxId;

  factory CanSub.fromJson(Map<String, dynamic> json) => CanSub(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        canSubscribe: json["canSubscribe"],
        boxId: json["BoxId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "canSubscribe": canSubscribe,
        "BoxId": boxId,
      };
}
