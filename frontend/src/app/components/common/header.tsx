export const Header = () => {
  return (
    <header className="w-full bg-gray-800 text-white p-4 shadow-md flex justify-between items-center">
      <div className="text-xl font-bold"><a href="/">Eitapi</a></div>
      <nav>
        <ul className="flex space-x-4">
          <li><a href="#" className="hover:underline">Home</a></li>
          <li><a href="#" className="hover:underline">Settings</a></li>
        </ul>
      </nav>
    </header>
  )
}