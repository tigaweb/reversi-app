export type HeaderProps = {
  title: string
}

const Header = ({ title }: HeaderProps) => {
  return (
    <header className="text-2xl font-bold h-12 flex justify-center items-center">
      <h1>{title}</h1>
    </header>
  );
};

export default Header;
