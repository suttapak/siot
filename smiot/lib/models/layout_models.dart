class Layout {
  Layout({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.i,
    required this.x,
    required this.y,
    required this.w,
    required this.h,
  });

  int id;
  DateTime createdAt;
  DateTime updatedAt;
  String i;
  int x;
  int y;
  int w;
  int h;

  factory Layout.fromJson(Map<String, dynamic> json) => Layout(
        id: json["id"],
        createdAt: DateTime.parse(json["createdAt"]),
        updatedAt: DateTime.parse(json["updatedAt"]),
        i: json["i"],
        x: json["x"],
        y: json["y"],
        w: json["w"],
        h: json["h"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "createdAt": createdAt.toIso8601String(),
        "updatedAt": updatedAt.toIso8601String(),
        "i": i,
        "x": x,
        "y": y,
        "w": w,
        "h": h,
      };
}
