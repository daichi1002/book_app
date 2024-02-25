import 'package:flutter/material.dart';
import 'package:flutter_app/model/books.dart';
import 'package:flutter_app/view/books/edit.dart';

// catsテーブルの中の1件のデータに対する操作を行うクラス
class BookDetail extends StatefulWidget {
  final int id;

  const BookDetail({Key? key, required this.id}) : super(key: key);

  @override
  _BookDetailState createState() => _BookDetailState();
}

class _BookDetailState extends State<BookDetail> {
  late Books books;
  bool isLoading = false;
  static const int textExpandedFlex = 1; // 見出しのexpaded flexの比率
  static const int dataExpandedFlex = 4; // 項目のexpanede flexの比率

  @override
  void initState() {
    super.initState();
    bookData();
  }

  Future bookData() async {
    setState(() => isLoading = true);
    books = Books(
      id: 1,
      title: "Book1",
      description: "Book1",
      score: 100,
    );
    setState(() => isLoading = false);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('book detail'),
        actions: [
          IconButton(
            onPressed: () async {
              await Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => BookEdit(
                    // 詳細更新画面を表示する
                    books: books,
                  ),
                ),
              );
              bookData(); // 更新後のデータを読み込む
            },
            icon: const Icon(Icons.edit), // 鉛筆マークのアイコンを表示
          ),
          IconButton(
            onPressed: () async {
              // ゴミ箱のアイコンが押されたときの処理を設定
              Navigator.of(context).pop(); // 削除後に前の画面に戻る
            },
            icon: const Icon(Icons.delete), // ゴミ箱マークのアイコンを表示
          )
        ],
      ),
      body: isLoading //「読み込み中」だったら「グルグル」が表示される
          ? const Center(
              child: CircularProgressIndicator(), // これが「グルグル」の処理
            )
          : Column(
              children: [
                Container(
                    // アイコンを表示
                    width: 80,
                    height: 80,
                    decoration: const BoxDecoration(
                        shape: BoxShape.circle, // 丸にする
                        image: DecorationImage(
                            fit: BoxFit.fill,
                            image: AssetImage('assets/icon/dora.png')))),
                Column(
                  // 縦並びで項目を表示
                  crossAxisAlignment: CrossAxisAlignment.stretch, // 子要素の高さを合わせる
                  children: [
                    Row(
                      children: [
                        const Expanded(
                          flex: textExpandedFlex,
                          child: Text(
                            'タイトル',
                            textAlign: TextAlign.center,
                          ),
                        ),
                        Expanded(
                          flex: dataExpandedFlex,
                          child: Container(
                            padding: const EdgeInsets.all(5.0),
                            child: Text(books.title),
                          ),
                        ),
                      ],
                    ),
                    Row(
                      children: [
                        const Expanded(
                          flex: textExpandedFlex,
                          child: Text(
                            '詳細',
                            textAlign: TextAlign.center,
                          ),
                        ),
                        Expanded(
                          flex: dataExpandedFlex,
                          child: Container(
                            padding: const EdgeInsets.all(5.0),
                            child: Text(books.description!),
                          ),
                        ),
                      ],
                    ),
                    Row(
                      children: [
                        const Expanded(
                          flex: textExpandedFlex,
                          child: Text(
                            '点数',
                            textAlign: TextAlign.center,
                          ),
                        ),
                        Expanded(
                          flex: dataExpandedFlex,
                          child: Container(
                            padding: const EdgeInsets.all(5.0),
                            child: Text(books.score.toString()),
                          ),
                        )
                      ],
                    ),
                  ],
                ),
              ],
            ),
    );
  }
}
