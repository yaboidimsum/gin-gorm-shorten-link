// import "./App.css";
import ShorternUrlForm from "./components/custom/ShortenUrlForm";
import ShortenUrlList from "./components/custom/ShortenUrlList";
export default function App() {
  return (
    <main className="min-h-screen w-screen flex justify-center items-center">
      <div className="flex w-full px-20 gap-4">
        <ShortenUrlList />
        <ShorternUrlForm />
      </div>
    </main>
  );
}
