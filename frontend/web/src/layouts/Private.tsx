import * as React from 'react';

type PrivateProps = { children: JSX.Element };

const Private = ({ children }: PrivateProps): JSX.Element => <>{children}</>;

export default Private;
