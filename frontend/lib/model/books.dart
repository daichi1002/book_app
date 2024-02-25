class Books {
  int? id;
  late String title;
  String? description;
  int? score;
  DateTime? createdAt;

  Books({
    this.id,
    required this.title,
    this.description,
    this.score,
    this.createdAt,
  });
}
