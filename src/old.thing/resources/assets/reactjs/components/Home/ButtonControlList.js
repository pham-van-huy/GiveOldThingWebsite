import React from 'react'
import classnames from 'classnames'

const AGrid = styled.a`
  /* Adapt the colors based on primary prop */
  &.active {
    background: #dcdbdb;
  }
`;

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import styled from 'styled-components'

const ButtonControlList = ({onClick, icon, active}) => {
    var classCon = classnames('control-item', { active: active });
    var color = '#c07e7e'
    if (active) {
        color = '#d6b6b6'
    }
    return (
        <AGrid className={classCon} onClick={onClick}>
            <FontAwesomeIcon icon={icon} size="2x" color={color}/>
        </AGrid>
    );
}

export default ButtonControlList