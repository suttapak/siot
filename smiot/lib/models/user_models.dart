import 'package:smiot/models/avatar_models.dart';
import 'package:smiot/models/role_models.dart';
import 'package:smiot/models/setting_models.dart';

class Users {
  Users({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.settingId,
    required this.avatar,
    required this.roles,
    required this.box,
    required this.setting,
  });

  String id;
  DateTime createdAt;
  DateTime updatedAt;
  String firstName;
  String lastName;
  String email;
  int settingId;
  Avatar avatar;
  List<Role> roles;
  List<Box> box;
  Setting setting;

  factory Users.fromJson(Map<String, dynamic> json) => Users(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        firstName: json["firstName"],
        lastName: json["lastName"],
        email: json["email"],
        settingId: json["settingId"],
        avatar: Avatar.fromJson(json["avatar"]),
        roles: List<Role>.from(json["roles"].map((x) => Role.fromJson(x))),
        box: List<Box>.from(json["box"].map((x) => Box.fromJson(x))),
        setting: Setting.fromJson(json["setting"]),
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "firstName": firstName,
        "lastName": lastName,
        "email": email,
        "settingId": settingId,
        "avatar": avatar.toJson(),
        "roles": List<dynamic>.from(roles.map((x) => x.toJson())),
        "box": List<dynamic>.from(box.map((x) => x.toJson())),
        "setting": setting.toJson(),
      };
}

class Box {
  Box({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.description,
    required this.ownerId,
  });

  String id;
  DateTime createdAt;
  DateTime updatedAt;
  String name;
  String description;
  String ownerId;

  factory Box.fromJson(Map<String, dynamic> json) => Box(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        description: json["description"],
        ownerId: json["ownerId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "name": name,
        "description": description,
        "ownerId": ownerId,
      };
}
