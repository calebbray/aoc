use anyhow::Result;
use hand::Hand;

mod hand {
    use std::{collections::HashMap, str::FromStr};

    #[derive(Debug, Clone, PartialEq)]
    pub struct Hand {
        cards: Vec<Card>,
        pub bid: usize,
        hand_hash: HashMap<usize, usize>,
    }

    impl Hand {
        pub fn compare(&self, other: &Hand, jacks_wild: Option<bool>) -> bool {
            if self.get_hand_rank(jacks_wild) == other.get_hand_rank(jacks_wild) {
                return self.get_higher_order(other, jacks_wild);
            }
            self.get_hand_rank(jacks_wild) < other.get_hand_rank(jacks_wild)
        }

        pub fn get_hand_rank(&self, jacks_wild: Option<bool>) -> Rank {
            if self.is_five_of_a_kind(jacks_wild) {
                return Rank::FiveKind;
            }
            if self.is_four_of_a_kind(jacks_wild) {
                return Rank::FourKind;
            }
            if self.is_full_house(jacks_wild) {
                return Rank::FullHouse;
            }
            if self.is_three_of_a_kind(jacks_wild) {
                return Rank::ThreeKind;
            }
            if self.is_two_pair() {
                return Rank::TwoPair;
            }
            if self.is_pair(jacks_wild) {
                return Rank::Pair;
            }

            Rank::HighCard
        }

        fn get_higher_order(&self, other: &Hand, jacks_wild: Option<bool>) -> bool {
            assert!(self.cards.len() == 5);
            for i in 0..self.cards.len() {
                if self.cards[i] == other.cards[i] {
                    continue;
                }
                let Card(mut me) = self.cards[i];
                if me == 11 && jacks_wild.is_some() {
                    me = 1;
                };
                let Card(mut o) = other.cards[i];
                if o == 0 && jacks_wild.is_some() {
                    o = 0;
                };
                return o < me;
            }
            unreachable!("There should always be cards in the hand")
        }

        fn is_five_of_a_kind(&self, jacks_wild: Option<bool>) -> bool {
            if !jacks_wild.is_some() {
                return self.hand_hash.keys().len() == 1;
            }
            let has_jacks = self.hand_hash.get(&11).is_some();
            return self.hand_hash.keys().len() == 1
                || self.hand_hash.keys().len() == 2 && has_jacks;
        }

        fn is_four_of_a_kind(&self, jacks_wild: Option<bool>) -> bool {
            if !jacks_wild.is_some() {
                return self.get_max_occuring() == 4;
            }

            let num_jacks = self.hand_hash.get(&11).unwrap_or(&0);
            *num_jacks + self.get_max_occuring() >= 4
        }

        fn is_full_house(&self, jacks_wild: Option<bool>) -> bool {
            let max_occuring = self.get_max_occuring();
            let is_normal_full_house = self.hand_hash.keys().len() == 2 && max_occuring == 3;
            if !jacks_wild.is_some() {
                return is_normal_full_house;
            }

            let num_jacks = self.hand_hash.get(&11).unwrap_or(&0);
            let is_two_pair = self.is_two_pair();

            is_normal_full_house || is_two_pair && num_jacks == &1
        }

        fn is_three_of_a_kind(&self, jacks_wild: Option<bool>) -> bool {
            if !jacks_wild.is_some() {
                return self.get_max_occuring() == 3;
            }

            let num_jacks = self.hand_hash.get(&11).unwrap_or(&0);
            *num_jacks + self.get_max_occuring() >= 3
        }

        fn is_two_pair(&self) -> bool {
            let mut pair_counts = 0;
            for (_k, v) in &self.hand_hash {
                if v == &2 {
                    pair_counts += 1;
                }
            }

            pair_counts == 2
        }

        fn is_pair(&self, jacks_wild: Option<bool>) -> bool {
            if !jacks_wild.is_some() {
                return self.get_max_occuring() == 2;
            }

            let num_jacks = self.hand_hash.get(&11).unwrap_or(&0);
            *num_jacks + self.get_max_occuring() >= 2
        }

        fn get_max_occuring(&self) -> usize {
            let mut max = 0;
            for (_k, v) in &self.hand_hash {
                if v > &max {
                    max = *v;
                }
            }
            max
        }
    }
    #[derive(Debug, PartialEq, PartialOrd, Eq, Ord)]
    pub enum Rank {
        FiveKind,
        FourKind,
        FullHouse,
        ThreeKind,
        TwoPair,
        Pair,
        HighCard,
    }

    impl FromStr for Hand {
        type Err = anyhow::Error;
        fn from_str(s: &str) -> Result<Self, Self::Err> {
            let (cards, bid) = s.split_once(" ").unwrap_or(("", ""));
            let cards: Vec<Card> = cards
                .chars()
                .map(|c| c.to_string().as_str().parse().unwrap())
                .collect();

            let bid = bid.parse()?;
            let mut disticts = HashMap::new();
            for Card(card) in &cards {
                disticts
                    .entry(*card)
                    .and_modify(|n| *n += 1)
                    .or_insert(1_usize);
            }

            Ok(Hand {
                cards,
                bid,
                hand_hash: disticts,
            })
        }
    }

    #[derive(Debug, Clone, Copy, PartialEq, Eq)]
    struct Card(usize);

    impl FromStr for Card {
        type Err = anyhow::Error;
        fn from_str(s: &str) -> Result<Self, Self::Err> {
            let value = match s.parse::<usize>() {
                Ok(x) => x,
                Err(_) => match s {
                    "T" => 10,
                    "J" => 11,
                    "Q" => 12,
                    "K" => 13,
                    "A" => 14,
                    _ => return Err(anyhow::anyhow!("invalid card encoding")),
                },
            };
            Ok(Self(value))
        }
    }
}

fn sort_hands(hands: &mut Vec<Hand>, jacks_wild: Option<bool>) {
    let len = hands.len();
    let mut swapped;

    loop {
        swapped = false;
        for i in 0..len - 1 {
            if hands[i].compare(&hands[i + 1], jacks_wild) {
                let tmp = hands[i].clone();
                hands[i] = hands[i + 1].clone();
                hands[i + 1] = tmp;
                swapped = true;
            }
        }

        if !swapped {
            break;
        }
    }
}

fn parse_hands(input: &str) -> Result<Vec<Hand>> {
    let mut hands = Vec::new();
    for line in input.lines() {
        hands.push(line.parse()?)
    }
    Ok(hands)
}

fn solve(hands: &mut Vec<Hand>, jacks_wild: Option<bool>) -> usize {
    let mut hands = hands.clone();
    sort_hands(&mut hands, jacks_wild);
    hands.iter().enumerate().map(|(i, h)| (i + 1) * h.bid).sum()
}

fn main() -> Result<()> {
    let mut hands = parse_hands(include_str!("./input.txt"))?;
    let part_one = solve(&mut hands, None);
    dbg!(part_one);
    let part_two = solve(&mut hands, Some(true));

    dbg!(part_two);
    Ok(())
}

#[cfg(test)]
mod tests {
    use crate::{
        hand::{Hand, Rank},
        sort_hands,
    };

    #[test]
    fn parses_a_hand() {
        let hand = "23456 123".parse::<Hand>();
        dbg!(&hand);
        assert!(hand.is_ok() == true)
    }

    #[test]
    fn evaluates_five_of_a_kind() {
        let (normal, with_wild) = (
            "KKKKK 123".parse::<Hand>().unwrap(),
            "KKKKJ 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::FiveKind);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::FiveKind);
        assert_eq!(with_wild.get_hand_rank(None), Rank::FourKind);
    }

    #[test]
    fn evaluates_four_of_a_kind() {
        let (normal, with_wild) = (
            "KKKTK 123".parse::<Hand>().unwrap(),
            "KKKTJ 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::FourKind);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::FourKind);
        assert_eq!(with_wild.get_hand_rank(None), Rank::ThreeKind);
    }

    #[test]
    fn evaluates_three_of_a_kind() {
        let (normal, with_wild) = (
            "KK4TK 123".parse::<Hand>().unwrap(),
            "KK4TJ 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::ThreeKind);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::ThreeKind);
        assert_eq!(with_wild.get_hand_rank(None), Rank::Pair);
    }

    #[test]
    fn evaluates_full_house() {
        let (normal, with_wild) = (
            "KK44K 123".parse::<Hand>().unwrap(),
            "KK44J 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::FullHouse);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::FullHouse);
        assert_eq!(with_wild.get_hand_rank(None), Rank::TwoPair);
    }

    #[test]
    fn evaluates_two_pair() {
        let (normal, with_wild) = (
            "KK445 123".parse::<Hand>().unwrap(),
            "KK42J 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::TwoPair);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::ThreeKind);
        assert_eq!(with_wild.get_hand_rank(None), Rank::Pair);
    }

    #[test]
    fn evaluates_pair() {
        let (normal, with_wild) = (
            "KK425 123".parse::<Hand>().unwrap(),
            "K942J 123".parse::<Hand>().unwrap(),
        );
        assert_eq!(normal.get_hand_rank(None), Rank::Pair);
        assert_eq!(with_wild.get_hand_rank(Some(true)), Rank::Pair);
        assert_eq!(with_wild.get_hand_rank(None), Rank::HighCard);
    }

    #[test]
    fn sorts_correctly() {
        let (normal, has_jack) = (
            "KK425 123".parse::<Hand>().unwrap(),
            "K922J 123".parse::<Hand>().unwrap(),
        );

        assert_eq!(normal.get_hand_rank(None), Rank::Pair);
        assert_eq!(has_jack.get_hand_rank(None), Rank::Pair);
        assert_eq!(has_jack.get_hand_rank(Some(true)), Rank::ThreeKind);

        assert!(normal.compare(&has_jack, None) == true);
        // Based on the fact that these are both one pair, the normal hand should be sorted to the
        // first spot since there is a king in the 1 index of the cards
        let mut hands = vec![has_jack.clone(), normal.clone()];
        sort_hands(&mut hands, Some(true));
        assert_eq!(hands[0], normal);
    }
}
