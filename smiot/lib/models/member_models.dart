class Member {
  Member({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.userAccessToken,
    required this.userId,
    required this.boxId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String userAccessToken;
  String userId;
  String boxId;

  factory Member.fromJson(Map<String, dynamic> json) => Member(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        userAccessToken: json["userAccessToken"],
        userId: json["userId"],
        boxId: json["BoxId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "userAccessToken": userAccessToken,
        "userId": userId,
        "BoxId": boxId,
      };
}
