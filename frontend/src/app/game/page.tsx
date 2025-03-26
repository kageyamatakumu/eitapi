"use client";

import { useState, useEffect } from 'react';
import { TypingGame } from '../components/game/typingGame';

export type Word = {
  englishWord: string;
  japaneseTranslation: string;
  pronunciation: string;
  exampleSentence: string;
};

export default function Home() {
  const englishWords: Word[] = [
    {
      englishWord: "apple",
      japaneseTranslation: "りんご",
      pronunciation: "æpl",
      exampleSentence: "This is an apple."
    },
    {
      englishWord: "banana",
      japaneseTranslation: "バナナ",
      pronunciation: "bəˈnɑːnə",
      exampleSentence: "I like to eat a banana."
    },
    {
      englishWord: "orange",
      japaneseTranslation: "オレンジ",
      pronunciation: "ˈɔːrɪndʒ",
      exampleSentence: "She peeled an orange."
    },
    {
      englishWord: "grape",
      japaneseTranslation: "ぶどう",
      pronunciation: "ɡreɪp",
      exampleSentence: "Grapes grow on vines."
    }
  ];

  const [currentWordIndex, setCurrentWordIndex] = useState<number>(0);
  const [currentPosition, setCurrentPosition] = useState<number>(0);
  const [isCompleted, setIsCompleted] = useState<boolean>(false);

  useEffect(() => {
    const handleKeyDown = async (event: KeyboardEvent) => {
      const currentWord = englishWords[currentWordIndex];
      if (event.key.toLowerCase() === currentWord.englishWord[currentPosition].toLowerCase()) {
        setCurrentPosition((prev) => prev + 1);

        if (currentPosition === currentWord.englishWord.length - 1) {
          if (currentWordIndex === englishWords.length - 1) {
            setIsCompleted(true);
          } else {
            setCurrentWordIndex((prev) => prev + 1);
            setCurrentPosition(0);
          }
        }
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [currentWordIndex, currentPosition]);

  if (isCompleted) {
    return (
      <main className="flex flex-col items-center justify-center h-screen">
        <div className="text-center w-full h-screen flex flex-col items-center justify-center">
          <div>
            <p className="text-black">Completed!</p>
          </div>
        </div>
      </main>
    );
  }

  return (
    <main className="flex flex-col items-center justify-center h-screen">
      <div className="text-center w-full h-screen flex flex-col items-center justify-center">
        <TypingGame englishWords={englishWords} currentWordIndex={currentWordIndex} currentPosition={currentPosition} />
      </div>
    </main>
  );
}