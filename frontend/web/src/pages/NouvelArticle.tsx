import * as React from 'react';
import { Button, HStack, Input, Textarea, VStack } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const NouvelArticle = (): JSX.Element => {
	const navigate = useNavigate();

	return (
		<VStack w="100%" h="100vh" p="64px" spacing="16px">
			<HStack w="100%" spacing="16px">
				<Input id="nouvel-article-title-input" variant="primary-1" placeholder="Titre du nouvel article" />
				<Input w="25%" id="nouvel-article-topic-input" variant="primary-1" placeholder="Sujet du nouvel article" />
			</HStack>
			<Textarea
				id="nouvel-article-content-textarea"
				variant="primary-1"
				placeholder="Contenu du nouvel article"
				flexGrow="2"
			/>
			<Button id="nouvel-article-publish-btn" variant="primary-1" onClick={() => navigate('/publications')}>
				Publier
			</Button>
			<HStack w="100%" spacing="16px">
				<Button id="nouvel-article-pre-visualize-btn" variant="secondary-4" isDisabled>
					Pré-visualisez votre article
				</Button>
				<Button id="nouvel-article-save-draft-btn" variant="secondary-1" onClick={() => navigate('/brouillons')}>
					Enregistrer dans les brouillons
				</Button>
			</HStack>
		</VStack>
	);
};

export default NouvelArticle;
