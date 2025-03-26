import {Header} from '../components/common/header';

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
    <div>
      <Header />
      <div className="flex h-screen flex-col md:flex-row md:overflow-hidden bg-white">
        <div className="flex-grow p-6 md:overflow-y-auto md:p-12">{children}</div>
      </div>
    </div>
  );
}