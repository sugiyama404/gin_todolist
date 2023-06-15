'use client';
import Link from 'next/link';
import UIkit from 'uikit'
import Icons from 'uikit/dist/js/uikit-icons'
import 'uikit/dist/css/uikit.css'
import 'uikit/dist/css/uikit.min.css'
UIkit.use(Icons)

export const Header = () => {
    return (
        <header>
            <nav className="uk-navbar-container">
                <div className="uk-container">
                    <div className="uk-navbar-left">
                        <ul className="uk-navbar-nav">
                            <li className="uk-active">
                                <Link href="/">Home</Link>
                            </li>
                            <li className="uk-active">
                                <Link href="/todo">Todo</Link>
                            </li>
                            <li className="uk-active">
                                <Link href="/page1">page1</Link>
                            </li>
                            <li className="uk-active">
                                <Link href="/page2">page2</Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </header>
    );
};
