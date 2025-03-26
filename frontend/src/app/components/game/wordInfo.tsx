"use client";

import React from 'react';
import { Word } from '../../game/page';

type WordInfoProps = {
  word: Word;
};

export const WordInfo: React.FC<WordInfoProps> = ({ word }) => {
  return (
    <div className="mt-4">
      <div className="text-black">{word.japaneseTranslation}</div>
      <div className="text-black">{word.pronunciation}</div>
      <div className="text-black">{word.exampleSentence}</div>
    </div>
  );
};
