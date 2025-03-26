"use client";

import React from 'react';
import { Word } from '../../game/page';

type WordDisplayProps = {
  word: Word;
  currentPosition: number;
};

export const WordDisplay: React.FC<WordDisplayProps> = ({ word, currentPosition }) => {
  return (
    <div>
      {word.englishWord.split('').map((char, index) => (
        <span
          key={index}
          className={index < currentPosition ? 'text-green-500' : 'text-black'}
        >
          {char}
        </span>
      ))}
    </div>
  );
};

