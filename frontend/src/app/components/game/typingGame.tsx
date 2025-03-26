import React from 'react';
import { WordDisplay } from './wordDisplay';
import { WordInfo } from './wordInfo';
import { Word } from '../../game/page';

type TypingGameProps = {
  englishWords: Word[];
  currentWordIndex: number;
  currentPosition: number;
};

export const TypingGame: React.FC<TypingGameProps> = ({ englishWords, currentWordIndex, currentPosition }) => {
  const currentWord = englishWords[currentWordIndex];

  return (
    <main className="flex flex-col items-center justify-center h-screen">
      <div className="text-center w-full h-screen flex flex-col items-center justify-center">
        <WordDisplay word={currentWord} currentPosition={currentPosition} />
        <WordInfo word={currentWord} />
      </div>
    </main>
  );
};