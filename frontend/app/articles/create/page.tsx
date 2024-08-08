"use client";
import { ArticleFormValues } from "@/app/types/schemas/article";
import ArticleForm from "@/components/ArticleForm";
import { useRouter } from "next/navigation";

const Page = () => {
  const router = useRouter();

  const onSubmit = async (value: ArticleFormValues) => {
    const { title, content } = value;
    try {
      await fetch(`${process.env.NEXT_PUBLIC_API_URL}/articles/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, content }),
      });
      router.push("/");
      router.refresh();
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <ArticleForm initValue={{ title: "", content: "" }} onSubmit={onSubmit} />
  );
};

export default Page;
