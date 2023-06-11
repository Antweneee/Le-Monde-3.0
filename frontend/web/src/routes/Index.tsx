import * as React from 'react';
import { BrowserRouter, Navigate, Route, Routes as RouterRoutes } from 'react-router-dom';

import AuthRoute from './Auth';
import HomePage from 'pages/Home';
import ConnexionPage from 'pages/Connexion';
import InscriptionPage from 'pages/Inscription';

const Routes = (): JSX.Element => (
	<BrowserRouter>
		<RouterRoutes>
			<Route path="/" element={<AuthRoute children={<HomePage />} />} />
			<Route path="/inscription" element={<AuthRoute children={<InscriptionPage />} />} />
			<Route path="/connexion" element={<AuthRoute children={<ConnexionPage />} />} />
			<Route path="*" element={<Navigate replace to="/" />} />
		</RouterRoutes>
	</BrowserRouter>
);

export default Routes;
