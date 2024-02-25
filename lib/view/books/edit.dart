import 'package:flutter/material.dart';
import 'package:flutter_app/model/books.dart';

class BookEdit extends StatefulWidget {
  final Books? books;

  const BookEdit({Key? key, this.books}) : super(key: key);

  @override
  _BookEditState createState() => _BookEditState();
}

const List<int> scores = <int>[1, 2, 3, 4, 5];

class _BookEditState extends State<BookEdit> {
  late int id;
  late String title;
  late String description;
  late int score;
  late DateTime createdAt;

  int isSelectedScore = scores.first;
  static const int textExpandedFlex = 1; // 見出しのexpaded flexの比率
  static const int dataExpandedFlex = 3; // 項

  @override
  void initState() {
    super.initState();
    id = widget.books?.id ?? 0;
    title = widget.books?.title ?? '';
    description = widget.books?.description ?? '';
    score = widget.books?.score ?? 0;
    createdAt = widget.books?.createdAt ?? DateTime.now();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("book edit"),
        actions: [
          buildSaveButton(),
        ],
      ),
      body: SingleChildScrollView(
        child: Column(children: <Widget>[
          Row(children: [
            const Expanded(
              flex: textExpandedFlex,
              child: Text(
                'タイトル',
                textAlign: TextAlign.center,
              ),
            ),
            Expanded(
              flex: dataExpandedFlex,
              child: TextFormField(
                maxLines: 1,
                initialValue: title,
                decoration: const InputDecoration(
                  hintText: 'タイトルを入力してください',
                ),
                validator: (title) => title != null && title.isEmpty
                    ? 'タイトルは必ず入れてね'
                    : null, // validateを設定
                onChanged: (title) => setState(() => this.title = title),
              ),
            ),
          ]),
          Row(children: [
            const Expanded(
              flex: textExpandedFlex,
              child: Text(
                '概要',
                textAlign: TextAlign.center,
              ),
            ),
            Expanded(
              flex: dataExpandedFlex,
              child: TextFormField(
                keyboardType: TextInputType.multiline,
                maxLines: null,
                initialValue: description,
                decoration: const InputDecoration(
                  hintText: '概要を入力してください',
                ),
                onChanged: (description) =>
                    setState(() => this.description = description),
              ),
            ),
          ]),
          Row(children: [
            const Expanded(
              flex: textExpandedFlex,
              child: Text(
                '点数',
                textAlign: TextAlign.center,
              ),
            ),
            Expanded(
              flex: dataExpandedFlex,
              child: DropdownButton(
                items: scores.map<DropdownMenuItem<int>>((int value) {
                  return DropdownMenuItem(
                    value: value,
                    child: Text(value.toString()),
                  );
                }).toList(),
                value: isSelectedScore,
                onChanged: (int? value) {
                  setState(() {
                    isSelectedScore = value!;
                  });
                },
              ),
            ),
          ]),
        ]),
      ),
    );
  }

  Widget buildSaveButton() {
    final isFormValid = title.isNotEmpty;

    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(
          foregroundColor: Colors.white,
          backgroundColor:
              isFormValid ? Colors.redAccent : Colors.grey.shade700,
        ),
        onPressed: createOrUpdateBook,
        child: const Text('保存'), // 保存ボタンを押したら実行する処理を指定する
      ),
    );
  }

  void createOrUpdateBook() async {
    final isUpdate = (widget.books != null); // 画面が空でなかったら

    if (isUpdate) {
      await updateBook(); // updateの処理
    } else {
      await createBook(); // insertの処理
    }

    Navigator.of(context).pop(); // 前の画面に戻る
  }

  // 更新処理の呼び出し
  Future updateBook() async {
    //   final cat = widget.books!.copy(
    //     // 画面の内容をcatにセット
    //     title: title,
    //     description: description,
    //   );

    //   await DbHelper.instance.update(cat); // catの内容で更新する
  }

  // 追加処理の呼び出し
  Future createBook() async {
    // final cat = Books(
    //   // 入力された内容をcatにセット
    //   title: title,
    //   description: description,
    // );
    // await DbHelper.instance.insert(cat); // catの内容で追加する
  }
}
