import { Article } from "@/app/types/types";
import ArticleCard from "./ArticleCard";

interface ArticlesProps {
  articles: Article[];
}

const ArticleCardList = ({ articles }: ArticlesProps) => {
  return (
    <div className="grid lg:grid-cols-3 px-4 py-4 gap-4">
      {articles.map((article: Article) => (
        <ArticleCard key={article.id} article={article} />
      ))}
    </div>
  );
};

export default ArticleCardList;
