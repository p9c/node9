package css

func ROOT() string {
	return `
:root {
--blue: #3030cf;
--light-blue: #30cfcf;

--green: #30cf30;
--orange: #ff7500;
--yellow: #ffd500;
--red: #cf3030;
--purple: #cf30cf;

--dark: #303030;
--white: #fcfcfc;

--dark-gray:#4d4d4d;
--gray:#808080;
--light-gray:#c0c0c0;

--light-green:#80cf80;
--light-orange:#cf9480;
--light-yellow:#cfcf80;
--light-red:#cf8080;
--light-purple:#cf80cf;

--dark-blue:#303080;
--dark-green:#308030;
--dark-orange:#804430;
--dark-yellow:#808030;
--dark-red:#803030;
--dark-purple:#803080;

--green-blue:#308080;
--light-green-blue:#80a8a8;
--dark-green-blue:#305858;
--green-orange:#80a830;
--light-green-orange:#a8bc80;
--dark-green-orange:#586c30;
--green-yellow:#80cf30;
--light-green-yellow:#a8cf80;
--dark-green-yellow:#588030;
--green-red:#808030;
--light-green-red:#a8a880;
--dark-green-red:#585830;
--blue-orange:#80a830;
--light-blue-orange:#a8bc80;
--dark-blue-orange:#583a58;
--bluered:#803080;
--light-blue-red:#a880a8;
--dark-blue-red:#583058;
--dark:#303030;
--dark-grayii:#424242;
--dark-grayi:#535353;
--dark-gray:#656565;
--gray:#808080;
--light-gray: #888888;
--light-grayi: #9a9a9a;
--light-grayii: #acacac;
--light-grayiii: #bdbdbd;
--light:#cfcfcf;



--border-light: rgba(255, 255, 255, .62);
--border-dark: rgba(0,0,0, .38);

--trans-light: rgba(255, 255, 255, .24);
--trans-dark: rgba(0,0,0, .24);
--trans-gray: rgba(48,48,48, .38);

--fonta:'Roboto';
--fontb:'Abril Fatface';
--fontc:'Oswald';
--big-title:'Vollkorn SC';

--base:var(--white);
--pri: var(--blue);
--sec: var(--light-blue);
--btn-tx: var(--base);
--btn-h-tx: #fff;
--btn-bg: var(--dark-green-blue);
--btn-h-bg: var(--green-blue);

--space-02: .12rem;
--space-05: .25rem;
--space-1: .5rem;  
--space-2: 1rem;   
--space-3: 1.5rem; 
--space-4: 2rem;   
--space-5: 2.5rem; 
--space-6: 3rem;   
--space-7: 3.5rem; 
--space-8: 4rem;   



--box-shadow-b: 0 1px 0 0 var(--black);
--box-shadow-l: 0 1px 0 0 var(--white);
--box-shadow-inset :inset 0 0 0 1px var(--sec);

--pri:var(--dark)
--sec:var(--light)
--ter:var(--gray)

--side:var(--)
}
`
}
