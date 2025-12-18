import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import {
  Card,
  // CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import React, { useState } from "react";

interface Form {
  title: string;
  original_url: string;
}

export default function ShorternUrlForm() {
  const [form, setForm] = useState<Form>({ title: "", original_url: "" });
  const [isLoading, setIsLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setForm((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!form.title && !form.original_url)
      return alert("There are empty fields!");

    setIsLoading(true);

    try {
      const response = await fetch("http://localhost:8000/v1/shorten", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(form),
      });

      if (!response.ok) {
        throw new Error("Failed to create link");
      }

      setForm({ title: "", original_url: "" });
    } catch (error) {
      console.error("Error submitting form: ", error);
      alert("something went wrong");
    } finally {
      setIsLoading(false);
      window.location.reload();
    }
  };

  return (
    <>
      <div className="w-1/4  ">
        <Card className="w-full">
          <CardHeader>
            <CardTitle>Shorten your link</CardTitle>
            <CardDescription>Enter your link down below</CardDescription>
          </CardHeader>
          <CardContent>
            <form>
              <div className="flex flex-col gap-6">
                <div className="grid gap-2">
                  <Label htmlFor="title">Title</Label>
                  <Input
                    id="title"
                    type="text"
                    placeholder="What's this about?"
                    name="title"
                    value={form.title}
                    onChange={handleChange}
                    disabled={isLoading}
                    required
                  />
                </div>
                <div className="grid gap-2">
                  <div className="flex items-center">
                    <Label htmlFor="original_url">Link</Label>
                  </div>
                  <Input
                    id="original_url"
                    type="text"
                    placeholder="Put the link here"
                    name="original_url"
                    value={form.original_url}
                    onChange={handleChange}
                    disabled={isLoading}
                    required
                  />
                </div>
              </div>
            </form>
          </CardContent>
          <CardFooter className="flex-col gap-2">
            <Button
              type="submit"
              className="w-full"
              disabled={isLoading}
              onClick={handleSubmit}
            >
              {isLoading ? "Loading" : "Shorten Link"}
            </Button>
          </CardFooter>
        </Card>
      </div>
    </>
  );
}
