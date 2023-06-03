import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <header className="flex items-center border-2 border-black h-[8vh] text-md fornt-bold justify-between px-5 shadow-md">
            <figure>
                <figcaption>Logo</figcaption>
            </figure>
            <nav className="flex justify-beetween w-[12vw]">
                <Link to="/add">Add Data</Link>
                <Link to="/view">View Data</Link>

            </nav>

        </header>
    );
};
export default Navbar;
