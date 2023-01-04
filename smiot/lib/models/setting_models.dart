import 'package:smiot/models/nofitication_models.dart';

class Setting {
  Setting({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.userId,
    required this.notification,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String userId;
  Notification notification;

  factory Setting.fromJson(Map<String, dynamic> json) => Setting(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        userId: json["userId"],
        notification: Notification.fromJson(json["notification"]),
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "userId": userId,
        "notification": notification.toJson(),
      };
}
