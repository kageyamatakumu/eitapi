import {Header} from "./components/common/header";

export default function Home() {
  const books = [
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "初級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "初級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "初級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "初級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "初級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "中級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "中級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "中級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "中級コース" },
    { title: "TOEIC L&R TEST 出る単特急 金のフレーズ", level: "中級コース" },
  ];

  return (
    <div className="bg-white flex flex-col min-h-screen">
      <Header />
      <div className="flex-grow px-4 py-4">
        <h1 className="text-3xl font-bold text-center text-gray-800 mb-8">
          Eitapi
        </h1>
        {/* 初級コース */}
        <div className="mb-8">
          <h2 className="text-2xl font-semibold text-gray-700 mb-4">
            初級コース
          </h2>
          <div className="flex flex-wrap justify-center">
            {books
              .filter((book) => book.level === "初級コース")
              .map((book, index) => (
                <div
                  key={index}
                  className="w-48 p-4 m-2 shadow-md rounded-lg bg-sky-100"
                >
                  <div className="h-32 bg-sky-200 rounded-t-lg">
                  </div>
                  <div className="p-2">
                    <p className="text-sm text-gray-600">{book.title}</p>
                  </div>
                </div>
              ))}
          </div>
        </div>

        {/* 中級コース */}
        <div>
          <h2 className="text-2xl font-semibold text-gray-700 mb-4">
            中級コース
          </h2>
          <div className="flex flex-wrap justify-center">
            {books
              .filter((book) => book.level === "中級コース")
              .map((book, index) => (
                <div
                  key={index}
                  className="w-48 p-4 m-2 shadow-md rounded-lg bg-sky-100"
                >
                  <div className="h-32 bg-sky-200 rounded-t-lg">
                  </div>
                  <div className="p-2">
                    <p className="text-sm text-gray-600">{book.title}</p>
                  </div>
                </div>
              ))}
          </div>
        </div>
      </div>
      {/* フッターなどの追加コンテンツ */}
      <div className="p-4 text-center text-gray-500">
        © 2024 Eitapi
      </div>
    </div>
  );
}