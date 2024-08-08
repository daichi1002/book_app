"use client";
import { Article } from "@/app/types/types";
import { Button } from "@/components/ui/button";
import { FilePenLine, Trash2 } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";

const getArticle = async (id: number) => {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/articles/${id}`,
    {
      cache: "no-store",
    }
  );

  const article: Article = await response.json();

  return article;
};

const Page = async ({ params }: { params: { id: number } }) => {
  const router = useRouter();
  const article = await getArticle(params.id);

  const { title, content, createdAt } = article;

  const deleteArticle = async () => {
    try {
      await fetch(`${process.env.NEXT_PUBLIC_API_URL}/articles/${params.id}`, {
        method: "DELETE",
      });
      router.push("/");
      router.refresh();
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className="mx-auto max-w-4xl p-4">
      <div className="mb-8">
        <h1 className="text-2xl font-bold">{title}</h1>
        <p className="text-gray-700">{createdAt.toString()}</p>
      </div>

      <div className="mb-8">
        <p className="text-gray-900">{content}</p>
      </div>

      <div className="flex space-x-4">
        <Link
          href={"/"}
          className="bg-blue-500 text-white font-bold py-2 px-4 rounded-md"
        >
          Back
        </Link>
        <Button variant="outline" size="icon">
          <FilePenLine className="h-4 w-4" onClick={deleteArticle} />
        </Button>
        <Button variant="outline" size="icon">
          <Trash2 className="h-4 w-4" onClick={deleteArticle} />
        </Button>
      </div>
    </div>
  );
};

export default Page;
