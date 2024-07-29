import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Link from "next/link";

const BBSCard = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Create project</CardTitle>
        <CardDescription>Deploy your new project in one-click.</CardDescription>
      </CardHeader>
      <CardContent>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Vel porro iusto
        aut. Dicta quam consectetur, itaque iure provident quibusdam illum fuga
        velit eos aut in quos id excepturi voluptates similique.
      </CardContent>
      <CardFooter className="flex justify-between">
        <Link href={"/bbs-posts/1"} className="text-blue-500">
          Read More
        </Link>
      </CardFooter>
    </Card>
  );
};

export default BBSCard;
