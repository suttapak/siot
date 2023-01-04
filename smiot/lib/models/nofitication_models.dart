class Notification {
  Notification({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.notificationState,
    required this.settingId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  bool notificationState;
  int settingId;

  factory Notification.fromJson(Map<String, dynamic> json) => Notification(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        notificationState: json["notificationState"],
        settingId: json["settingId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "notificationState": notificationState,
        "settingId": settingId,
      };
}
