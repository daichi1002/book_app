import 'package:flutter/material.dart';
import 'package:flutter_app/model/books.dart';
import 'package:flutter_app/view/books/edit.dart';

class BookList extends StatefulWidget {
  const BookList({Key? key}) : super(key: key);

  @override
  _BookListPageState createState() => _BookListPageState();
}

class _BookListPageState extends State<BookList> {
  List<Books> bookList = [
    Books(
      id: 1,
      title: "Book1",
      description: "Book1",
      score: 100,
    ),
    Books(
      id: 2,
      title: "Book2",
      description: "Book2",
      score: 100,
    ),
  ];
  bool isLoading = false;

  @override
  void initState() {
    super.initState();
    getBooksList();
  }

  Future getBooksList() async {
    setState(() {
      isLoading = true;
    });
    // bookを取得する処理
    setState(() {
      isLoading = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text("books")),
      body: isLoading
          ? const Center(
              child: CircularProgressIndicator(),
            )
          : SizedBox(
              child: ListView.builder(
                itemCount: bookList.length,
                itemBuilder: (BuildContext context, int index) {
                  final book = bookList[index];
                  return Card(
                    child: InkWell(
                      child: Padding(
                        padding: const EdgeInsets.all(15.0),
                        child: Row(
                          children: <Widget>[
                            Container(
                              width: 80,
                              height: 80,
                              decoration: const BoxDecoration(
                                  shape: BoxShape.circle,
                                  image: DecorationImage(
                                      fit: BoxFit.fill, image: AssetImage(''))),
                            ),
                            Text(book.title,
                                style: const TextStyle(fontSize: 30)),
                          ],
                        ),
                      ),
                      // onTap: () async {
                      //   await Navigator.of(context).push(
                      //     MaterialPageRoute(
                      //         builder: (context) =>
                      //             BookDetail(id: book.id!)),
                      //   );
                      //   getBooksList();
                      // },
                    ),
                  );
                },
              ),
            ),
      floatingActionButton: FloatingActionButton(
        child: const Icon(Icons.add),
        onPressed: () async {
          await Navigator.of(context)
              .push(MaterialPageRoute(builder: (context) => const BookEdit()));
          getBooksList();
        },
      ),
    );
  }
}
