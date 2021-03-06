﻿using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Title
{
    public class SignUpView : MonoBehaviour, IView
    {
        [SerializeField]
        Button signUpButton;
        [SerializeField]
        InputField nameInputField;

        public string InputtedName { get { return nameInputField.text; } }

        public IObservable<Unit> SignUpButtonClickedAsObservable()
        {
            return signUpButton.OnSingleClickAsObservable();
        }

        public IObservable<string> NameChangedAsObservable()
        {
            return nameInputField.OnValueChangedAsObservable();
        }

        public void Show()
        {
            gameObject.SetActive(true);
        }

        public void Hide()
        {
            gameObject.SetActive(false);
        }

        public void FocusToNameInputField()
        {
            nameInputField.ActivateInputField();
            nameInputField.Select();
        }
    }
}
