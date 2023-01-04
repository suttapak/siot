class Avatar {
  Avatar({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.title,
    required this.url,
    required this.userId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String title;
  String url;
  String userId;

  factory Avatar.fromJson(Map<String, dynamic> json) => Avatar(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        title: json["title"],
        url: json["url"],
        userId: json["userId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "title": title,
        "url": url,
        "userId": userId,
      };
}
