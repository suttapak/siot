class BoxMemberPermission {
  BoxMemberPermission({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.canRead,
    required this.canWrite,
    required this.boxMemberId,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  bool canRead;
  bool canWrite;
  int boxMemberId;

  factory BoxMemberPermission.fromJson(Map<String, dynamic> json) =>
      BoxMemberPermission(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        canRead: json["canRead"],
        canWrite: json["canWrite"],
        boxMemberId: json["boxMemberId"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "canRead": canRead,
        "canWrite": canWrite,
        "boxMemberId": boxMemberId,
      };
}
