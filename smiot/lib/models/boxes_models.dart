// To parse this JSON data, do
//
//     final boxes = boxesFromJson(jsonString);

import 'package:meta/meta.dart';
import 'package:smiot/models/box_secret_models.dart';
import 'package:smiot/models/can_pub_models.dart';
import 'package:smiot/models/can_sub_models.dart';
import 'dart:convert';

import 'package:smiot/models/member_models.dart';

List<Boxes> boxesFromJson(String str) =>
    List<Boxes>.from(json.decode(str).map((x) => Boxes.fromJson(x)));

String boxesToJson(List<Boxes> data) =>
    json.encode(List<dynamic>.from(data.map((x) => x.toJson())));

class Boxes {
  Boxes({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.name,
    required this.description,
    required this.ownerId,
    required this.members,
    required this.boxSecret,
    required this.canSub,
    required this.canPub,
  });

  String id;
  DateTime createdAt;
  DateTime updatedAt;
  String name;
  String description;
  String ownerId;
  List<Member> members;
  BoxSecret boxSecret;
  CanSub canSub;
  CanPub canPub;

  factory Boxes.fromJson(Map<String, dynamic> json) => Boxes(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        name: json["name"],
        description: json["description"],
        ownerId: json["ownerId"],
        members:
            List<Member>.from(json["members"].map((x) => Member.fromJson(x))),
        boxSecret: BoxSecret.fromJson(json["boxSecret"]),
        canSub: CanSub.fromJson(json["canSub"]),
        canPub: CanPub.fromJson(json["canPub"]),
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "name": name,
        "description": description,
        "ownerId": ownerId,
        "members": List<dynamic>.from(members.map((x) => x.toJson())),
        "boxSecret": boxSecret.toJson(),
        "canSub": canSub.toJson(),
        "canPub": canPub.toJson(),
      };
}
