import Link from "next/link";

const Nav = () => {
  return (
    <nav className="bg-customBlue">
      <ul>
        <li>
          <a href="/">Home</a>
        </li>
        <li>
          <Link href="/signup">SignUp</Link>
        </li>
        <li>
          <Link href="/login">Login</Link>
        </li>
      </ul>
    </nav>
  );
};

export default Nav;
