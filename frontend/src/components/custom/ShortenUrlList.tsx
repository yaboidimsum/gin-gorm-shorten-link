import { useState, useEffect } from "react";

interface Link {
  id: number;
  title: string;
  original_url: string;
  short_code: string;
}

export default function ShortenUrlList() {
  const [links, setLinks] = useState<Link[]>([]);

  useEffect(() => {
    const data = fetch("http://localhost:8000/v1/shortlink")
      .then((res) => res.json())
      .then((jsonData) => {
        setLinks(jsonData.data);
      })
      .catch((err) => console.log("Failed to fetch data:", err));
  }, []);

  console.log(links);

  return (
    <>
      <div className="w-3/4 ">
        {/* <span>This is Card List</span> */}
        <ul className="grid grid-cols-3 gap-4">
          {links &&
            links.map((item) => (
              <li
                key={item.id}
                className="group flex flex-wrapped flex-col gap-3 rounded-md border border-slate-200 bg-white p-4 shadow-sm transition-all hover:shadow-md"
              >
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-3">
                    {/* ID Badge - Style ala 'Badge' shadcn */}
                    <span className="flex h-6 w-6 text-md items-center justify-center rounded-full bg-slate-100 text-xs font-medium text-slate-600">
                      {item.id}
                    </span>
                    {/* Title - Typography tebal dan rapi */}
                    <h3 className="font-semibold text-md leading-none tracking-tight text-slate-900">
                      {item.title}
                    </h3>
                  </div>
                </div>

                {/* Description - Warna text-muted-foreground (abu-abu) */}
                <div className="flex gap-8">
                  <div className="flex flex-col">
                    <span className="font-semibold text-sm">Original Url</span>
                    <p className="text-xs text-slate-500 line-clamp-2 break-all">
                      {item.original_url}
                    </p>
                  </div>
                  <div className="flex flex-col">
                    <span className="font-semibold text-sm">Shorten Url</span>
                    <a
                      href={`http://localhost:8000/${item.short_code}`}
                      target="_blank"
                      className="text-xs text-slate-500 line-clamp-2 break-all"
                    >
                      {`http://localhost:8000/${item.short_code}`}
                    </a>
                  </div>
                </div>
              </li>
            ))}
        </ul>
      </div>
    </>
  );
}
