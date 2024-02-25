import 'package:flutter/material.dart';

import 'view/books/list.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      //初期画面の設定
      title: 'books',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      routes: <String, WidgetBuilder>{
        '/': (_) => const BookList(), //cat_list.dartを呼び出し
      },
    );
  }
}
