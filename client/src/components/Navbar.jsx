const Navbar = () => {
  return (
    <nav>
      <a className="logo" href="#">
        <img></img>
      </a>
      <a className="title" href="#">
        Tutorixy
      </a>
      <ul>
        <li>
          <a href="/test-prep">Test Prep</a>
        </li>
        <li>
          <a href="/tutoring">Tutoring</a>
        </li>
        <li>
          <a href="/about">About</a>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
