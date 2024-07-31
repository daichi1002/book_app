import ArticleCardList from "@/components/ArticleCardList";
import { Article } from "./types/types";

const getArticles = async () => {
  const response = await fetch("http://localhost:8080/articles", {
    cache: "no-store",
  });

  const articles: Article[] = await response.json();

  return articles;
};

export default async function Home() {
  const articles = await getArticles();
  return (
    <main>
      <ArticleCardList articles={articles} />
    </main>
  );
}
