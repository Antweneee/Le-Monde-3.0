import * as React from 'react';
import { BrowserRouter, Navigate, Route, Routes as RouterRoutes } from 'react-router-dom';

import AuthRoute from './Auth';
import PrivateRoute from './Private';
import HomePage from 'pages/Home';
import ConnexionPage from 'pages/Connexion';
import InscriptionPage from 'pages/Inscription';
import Favoris from 'pages/Favoris';
import MarquePages from 'pages/MarquePages';
import Nouveautes from 'pages/Nouveautes';
import Decouvertes from 'pages/Decouvertes';
import NouvelArticle from 'pages/NouvelArticle';
import Publications from 'pages/Publications';
import Brouillons from 'pages/Brouillons';
import Statistiques from 'pages/Statistiques';
import Reglages from 'pages/Reglages';
import Deconnexion from 'pages/Deconnexion';

const Routes = (): JSX.Element => (
	<BrowserRouter>
		<RouterRoutes>
			<Route path="/" element={<AuthRoute />}>
				<Route path="/" element={<HomePage />} />
				<Route path="/inscription" element={<InscriptionPage />} />
				<Route path="/connexion" element={<ConnexionPage />} />
			</Route>
			<Route path="/" element={<PrivateRoute />}>
				<Route path="/favoris" element={<Favoris />} />
				<Route path="/marque-pages" element={<MarquePages />} />
				<Route path="/nouveautes" element={<Nouveautes />} />
				<Route path="/decouvertes" element={<Decouvertes />} />
				<Route path="/nouvel-article" element={<NouvelArticle />} />
				<Route path="/publications" element={<Publications />} />
				<Route path="/brouillons" element={<Brouillons />} />
				<Route path="/statistiques" element={<Statistiques />} />
				<Route path="/reglages" element={<Reglages />} />
				<Route path="/deconnexion" element={<Deconnexion />} />
			</Route>
			<Route path="*" element={<Navigate replace to="/" />} />
		</RouterRoutes>
	</BrowserRouter>
);

export default Routes;
