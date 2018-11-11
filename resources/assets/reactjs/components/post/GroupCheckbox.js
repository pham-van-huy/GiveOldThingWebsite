import React from 'react'
import classnames from 'classnames'

const DivCheckbox = styled.a`
  /* Adapt the colors based on primary prop */
  .styled-input--rounded {
    label {
      &:before {
        border-radius: 10px;
      }
      &:after {
        border-radius: 6px;
      }
    }
  }
  .styled-input-single {
    position: relative;
    padding: 20px 10px 20px 40px;
    text-align: left;
    display: inline-block;
    label {
        font-size: 15px;
      cursor: pointer;
      &:before,
      &:after {
        content: '';
        position: absolute;
        top: 50%;
        border-radius: 50%;
      }
      &:before {
        left: 0;
        width: 30px;
        height: 30px;
        margin: -15px 0 0;
        background: #f7f7f7;
        box-shadow: 0 0 1px grey;
      }
      &:after {
        left: 5px;
        width: 20px;
        height: 20px;
        margin: -10px 0 0;
        opacity: 0;
        background: #37b2b2;
        transform: translate3d(-40px, 0, 0) scale(0.5);
        transition: opacity 0.25s ease-in-out, transform 0.25s ease-in-out;
      }
    }
    
    input[type="radio"],
    input[type="checkbox"] {
      position: absolute;
      top: 0;
      left: -9999px;
      visibility: hidden;
      
      &:checked + label {
        &:before {
        }
        &:after {
          transform: translate3d(0, 0, 0);
          opacity: 1;
        }
      }
    }
  }
`;

import styled from 'styled-components'

const GroupCheckbox = ({ onCheck, data }) => {
    var listCheckbox = data.map((item, key) => {
        var classN = 'fieldset-' + item.value
        var idN = 'radio-item-' + item.value
        var name = '' + Date.now() / 1000 | 0
        return (
            <div className="styled-input-single" key={key}>
                <input onChange={onCheck} type="radio" name={'checkbox_tt' + name} className={classN} id={idN} value={item.value} />
                <label htmlFor={idN}>{item.label}</label>
            </div>
        )
    })
    return (
        <DivCheckbox className="styled-input-container styled-input--rounded">
            {listCheckbox}
        </DivCheckbox>
    );
}

export default GroupCheckbox