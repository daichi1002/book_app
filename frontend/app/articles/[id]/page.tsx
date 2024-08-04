import { Article } from "@/app/types/types";
import Link from "next/link";

const getArticle = async (id: number) => {
  const response = await fetch(`http://localhost:8080/articles/${id}`, {
    cache: "no-store",
  });

  const article: Article = await response.json();

  return article;
};

const page = async ({ params }: { params: { id: number } }) => {
  const article = await getArticle(params.id);

  const { title, content, createdAt } = article;
  return (
    <div className="mx-auto max-w-4xl p-4">
      <div className="mb-8">
        <h1 className="text-2xl font-bold">{title}</h1>
        <p className="text-gray-700">{createdAt.toString()}</p>
      </div>

      <div className="mb-8">
        <p className="text-gray-900">{content}</p>
      </div>

      <Link
        href={"/"}
        className="bg-blue-500 text-white font-bold py-2 px-4 rounded-md"
      >
        Back
      </Link>
    </div>
  );
};

export default page;
