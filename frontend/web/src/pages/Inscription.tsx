import * as React from 'react';
import { Button, Input, Link } from '@chakra-ui/react';
import { Link as RouteLink, useNavigate } from 'react-router-dom';

const Inscription = (): JSX.Element => {
	const navigate = useNavigate();

	return (
		<>
			<Input id="inscription-email-input" variant="primary-1" placeholder="e-mail" />
			<Input id="inscription-username-input" variant="primary-1" placeholder="nom d'utilisateur" />
			<Input id="inscription-pwd-input" variant="primary-1" placeholder="mot de passe" type="password" />
			<Input
				id="inscription-confirmed-pwd-input"
				variant="primary-1"
				placeholder="confirmation du mot de passe"
				type="password"
			/>
			<Button id="inscription-inscription-btn" variant="primary-1" onClick={() => navigate('/favoris')}>
				Inscription
			</Button>
			<Link as={RouteLink} to="/connexion" w="100%">
				<Button id="inscription-connexion-btn" variant="secondary-4">
					Connexion
				</Button>
			</Link>
		</>
	);
};

export default Inscription;
