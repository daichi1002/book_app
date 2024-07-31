import { Article } from "@/app/types/types";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Link from "next/link";

interface ArticleProps {
  article: Article;
}

const ArticleCard = ({ article }: ArticleProps) => {
  const { id, title, content, createdAt } = article;
  return (
    <Card>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        {/* <CardDescription>{createdAt}</CardDescription> */}
      </CardHeader>
      <CardContent>{content}</CardContent>
      <CardFooter className="flex justify-between">
        <Link href={`/bbs-posts/${id}`} className="text-blue-500">
          Read More
        </Link>
      </CardFooter>
    </Card>
  );
};

export default ArticleCard;
